const ws = new WebSocket("ws://localhost:8080")

const messageInput = document.getElementById("messageInput") as HTMLInputElement;
const submitButton = document.getElementById("submitButton") as HTMLButtonElement;
const messageContainer = document.querySelector(".message-box") as HTMLDivElement;

const contact = document.querySelectorAll(".contact")

contact.forEach((c) => {
	c.addEventListener("click", (_) => {
		updateUI({user: "sender", text: "clicked"})
	})
})


type Message = {
	user: "sender" | "receiver"
	text: string
}



ws.addEventListener("message", (msg: MessageEvent) => {
    if (Array.isArray(msg.data)) {
		if (msg.data.length <= 0) return
        for (const buffer of msg.data) {
            if (buffer.isBuffer(buffer)) {
                const text = buffer.toString('utf-8');
                updateUI({ user: "receiver", text });
            }
        }
    } else if (msg.data.length > 0) {
			msg.data.forEach((message: Message) => {	
        		updateUI({ user: "receiver", text: message.text});
			})
    }

});




// update the ui by adding message to the screen
const updateUI = (msg: Message) => {
    const message = document.createElement("div");

    message.classList.add("message", msg.user);
    message.textContent = msg.text;

    messageContainer.appendChild(message);
	messageContainer.scrollTop = messageContainer.scrollHeight
};


// Clears the message input field
const clearInputField = () => {
	messageInput.value = ""
}

// sends the message after user clicks send message
const submitHandler = (event: Event) => {

    event.preventDefault();
    const inputElement = messageInput?.value;
	// check if the message input element exist 
    if (inputElement) {
            updateUI({user: "sender", text: inputElement});
			ws.send(JSON.stringify(inputElement));
			// ws.send(inputElement)
			clearInputField();
        }
    
};

// Event for the send message button
submitButton.addEventListener("click", submitHandler);

export {};
