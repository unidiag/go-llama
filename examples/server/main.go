package main

import (
	"fmt"
	"log"
	"net/http"

	llama "github.com/unidiag/go-llama"
)

var port = ":8080"

func main() {

	client := llama.New("http://192.168.1.55:80", "MAGICKSECRETKEY")
	//client.SetDefaults(0.5, 100)

	// Serve HTML form
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		html := `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>LLM Stream</title>
</head>
<body>
<h2>LLM Streaming Server</h2>

<form id="chatForm">
System:<br>
<textarea name="system" rows="3" cols="80">Ты профессиональный редактор EPG. Отвечай кратко.</textarea><br><br>

User:<br>
<textarea name="user" rows="3" cols="80"></textarea><br><br>

<button type="submit">Send</button>
</form>

<hr>

<pre id="output" style="white-space: pre-wrap;"></pre>

<script>
document.getElementById("chatForm").onsubmit = function(e) {
	e.preventDefault();

	const system = this.system.value;
	const user = this.user.value;

	const output = document.getElementById("output");
	output.textContent = "";

	const evtSource = new EventSource(
		"/stream?system=" + encodeURIComponent(system) +
		"&user=" + encodeURIComponent(user)
	);

	evtSource.onmessage = function(event) {
		output.textContent += event.data;
	};

	evtSource.onerror = function() {
		evtSource.close();
	};
};
</script>

</body>
</html>
`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	// Streaming endpoint
	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {

		system := r.URL.Query().Get("system")
		user := r.URL.Query().Get("user")

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		req := llama.ChatRequest{
			Messages: []llama.Message{
				{Role: "system", Content: system},
				{Role: "user", Content: user},
			},
		}

		err := client.ChatStream(req, func(token string) {
			fmt.Fprintf(w, "data: %s\n\n", token)
			flusher.Flush()
		})

		if err != nil {
			fmt.Fprintf(w, "data: ERROR: %v\n\n", err)
			flusher.Flush()
			return
		}

	})

	log.Println("Server started at " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
