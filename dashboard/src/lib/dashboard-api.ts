import { apiGet } from "./api-client";

export function getStats() {
  return apiGet("/api/dashboard/stats");
}

export function getHealth() {
  return apiGet("/api/dashboard/health");
}

export function getTraffic() {
  return apiGet("/api/dashboard/traffic");
}

export function getDashboardNodes() {
  return apiGet("/api/dashboard/nodes");
}
