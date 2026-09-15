const API_URL =
  process.env.NEXT_PUBLIC_KINGDNS_API || "http://localhost:8080";

export async function apiGet(path: string) {
  const token =
    typeof window !== "undefined"
      ? localStorage.getItem("kingdns_token")
      : null;

  const response = await fetch(`${API_URL}${path}`, {
    headers: {
      Authorization: token ? `Bearer ${token}` : "",
    },
  });

  return response.json();
}

export async function getDashboardStats() {
  return apiGet("/api/dashboard/stats");
}

export async function getNodes() {
  return apiGet("/api/dashboard/nodes");
}
