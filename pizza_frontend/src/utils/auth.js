export function getUserRole() {
  try {
    const token = localStorage.getItem("token");
    if (!token) return null;

    const base64Payload = token.split(".")[1];
    const payload = JSON.parse(atob(base64Payload));
    return payload.role || null;
  } catch {
    return null;
  }
}

export function isAdmin() {
  return getUserRole() === "admin";
}