import { useEffect, useState } from "react";
import { getDashboardStats } from "../lib/api-client";

export function useDashboard() {
  const [stats, setStats] = useState(null);

  useEffect(() => {
    getDashboardStats()
      .then(setStats)
      .catch(() => setStats(null));
  }, []);

  return stats;
}
