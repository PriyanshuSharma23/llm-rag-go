# RAG Application
This application utilizes ChromaDB and OpenAI to implement a Retrieval-Augmented Generation (RAG) system.

## Technologies Used:
1. Golang
2. ChiRouter
3. ChromaDB
4. OpenAI

## Running the Application Locally

You will need an OpenAI API Key to run this application. You can obtain your own key from the [OpenAI Dashboard](https://platform.openai.com/api-keys) or through Azure OpenAI. Set the `OPENAI_API_KEY` environment variable with your API key.

### Using Docker Compose (Recommended)
1. Clone the Git repository.
2. Run the following command:
   ```sh
   $ OPENAI_API_KEY=sk-** docker compose up
   ```

### Running Locally Without Docker
1. Start a ChromaDB server. Follow the instructions in the [ChromaDB documentation](https://docs.trychroma.com/guides#running-chroma-in-client-server-mode) to run it in client-server mode.
2. Set the `CHROMA_URL` environment variable to the URL of your ChromaDB instance from Step 1.
3. Clone this repository.
4. Run the application with the following command:
   ```sh
   CHROMA_URL=... OPENAI_API_KEY=... go run ./cmd/api/
   ```
