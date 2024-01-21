export const signup = async (name: string, password: string) => {
  try {
    const response = await fetch("http://44.221.67.84:8080/signUp", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username: name, password: password }),
    });

    if (!response.ok) {
      throw new Error("Login failed");
    }

    const data = await response.json();
    return data.token;
  } catch (error) {
    console.error("Login error:", error);
    throw error;
  }
};
