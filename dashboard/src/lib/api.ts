const API_URL =
  process.env.NEXT_PUBLIC_KINGDNS_API || "http://localhost:8080";

export async function dashboardRequest(path: string) {
  const response = await fetch(`${API_URL}${path}`);
  return response.json();
}

export async function getStats() {
  return dashboardRequest("/api/dashboard/stats");
}
