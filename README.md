# 🚀 Go Web Server with Language Model Poem Generator

This repository contains a simple Go web server that serves an HTML page and provides an API endpoint for generating poems using a language model.

## Features

- **Web Server**: The Go server listens on port 8080 and serves an HTML page at the root ("/") path. The HTML page can be customized to display relevant content.

- **Poem Generation API**: The server exposes a `/generate` API endpoint that accepts a prompt (in this case, a poem title) and generates a creative poem using a language model. If the title is empty, an error is returned.

## Getting Started

1. **Clone this repository**:
   ```bash
   git clone https://github.com/yourusername/go-poem-generator.git
   cd go-poem-generator
   ```

2. **Set up your API key for the language model**:
   - Obtain GEMINI API key for the language model.
   - Set the API key as an environment variable:
     ```bash
     export API_KEY=your_api_key_here
     ```

3. **Build and run the Go server**:
   ```bash
   go build
   ./go-poem-generator
   ```

4. **Access the web page**:
   Open your browser and navigate to `http://localhost:8080`.

## API Usage

- **Generate a Poem**:
  - **Endpoint**: `POST /generate`
  - **Request Body**:
    ```json
    {
      "prompt": "Your Poem Title"
    }
    ```
  - **Response**:
    ```json
    {
      "poem": "Generated poem goes here..."
    }
    ```

## Customization

Feel free to customize the HTML page, add more endpoints, or enhance the functionality as needed for your project.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
