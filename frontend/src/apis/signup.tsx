export const signup = async (name: string, password: string) => {
  try {
    const response = await fetch('http://localhost:8080/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username: name, password: password }),
    });

    if (!response.ok) {
      throw new Error('Login failed');
    }

    const data = await response.json();
    return data; // contains token or user details
  } catch (error) {
    console.error('Login error:', error);
    throw error;
  }
};