export const getChat = async (token: string, otherPersonId: number) => {
  try {
    const response = await fetch(
      `http://44.221.67.84:8080/chatMessages/${otherPersonId}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: null,
      },
    );

    if (!response.ok) {
      throw new Error("Message sending failed");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Get Messages error:", error);
    throw error;
  }
};
