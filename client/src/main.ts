document.addEventListener("DOMContentLoaded", () => {
  const messageInput = document.getElementById("messageInput") as HTMLInputElement;
  const sendButton = document.getElementById("sendButton");

  if (sendButton && messageInput) {
    sendButton.addEventListener("click", () => {
      const messageText = messageInput.value.trim();
      if (messageText !== "") {
        // Create a new message element
        const message = document.createElement("div");
        message.className = "message sender"; // You can adjust the class here
        message.innerHTML = `<p>${messageText}</p>`;

        // Append the new message to the conversation
        const convo = document.querySelector(".convo");
        if (convo) {
          convo.appendChild(message);

          // Clear the input field
          messageInput.value = "";

          // Automatically scroll to the bottom to show the latest message
          convo.scrollTop = convo.scrollHeight;
        }
      }
    });
  }
});
