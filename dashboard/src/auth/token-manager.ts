export function setToken(token: string) {
  localStorage.setItem("kingdns_token", token);
}

export function clearToken() {
  localStorage.removeItem("kingdns_token");
}

export function getToken() {
  return localStorage.getItem("kingdns_token");
}
