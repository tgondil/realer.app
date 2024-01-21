export const sendMessage = async (
  token: string,
  receiverId: number,
  messageContent: string,
) => {
  try {
    const response = await fetch("http://44.221.67.84:8080/sendMessage", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ id: receiverId, content: messageContent }),
    });

    if (!response.ok) {
      throw new Error("Message sending failed");
    }

    const data = await response.json();
  } catch (error) {
    console.error("SendMessage error:", error);
    throw error;
  }
};
