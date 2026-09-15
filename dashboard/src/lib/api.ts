const API_URL =
  process.env.NEXT_PUBLIC_KINGDNS_API || "http://localhost:8080";

export async function getDashboardStats() {
  const response = await fetch(`${API_URL}/api/dashboard/stats`);
  return response.json();
}
