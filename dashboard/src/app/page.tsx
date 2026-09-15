import MetricsCard from "../components/metrics-card";

export default function DashboardPage() {
  return (
    <main>
      <h1>KingDNS Control Center</h1>

      <MetricsCard title="Nodes Online" value="Loading..." />
      <MetricsCard title="DNS Queries" value="Loading..." />
      <MetricsCard title="Latency" value="Loading..." />
      <MetricsCard title="Health" value="Loading..." />
    </main>
  );
}
