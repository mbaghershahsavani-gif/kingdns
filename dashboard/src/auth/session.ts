export function saveToken(token: string) {
  if (typeof window !== "undefined") {
    localStorage.setItem("kingdns_token", token);
  }
}

export function getToken() {
  if (typeof window !== "undefined") {
    return localStorage.getItem("kingdns_token");
  }

  return null;
}
