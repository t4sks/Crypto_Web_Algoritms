async function sendRequest() {
    const text = document.getElementById("text").value;
    const key = parseInt(document.getElementById("key").value);
    const operation = document.getElementById("operation").value;

    const resultBlock = document.getElementById("result");
    resultBlock.textContent = "Обработка...";

    try {
        const response = await fetch("/api/scytale", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                text: text,
                key: key,
                operation: operation
            })
        });

        const data = await response.json();

        if (data.result) {
            resultBlock.textContent = data.result;
        } else if (data.error) {
            resultBlock.textContent = "Ошибка: " + data.error;
        } else {
            resultBlock.textContent = "Неизвестный ответ сервера";
        }

    } catch (err) {
        resultBlock.textContent = "Ошибка соединения";
    }
}