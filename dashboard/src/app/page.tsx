import Header from "../components/header";
import StatCard from "../components/stat-card";
import NodeTable from "../components/node-table";
import HealthChart from "../components/health-chart";
import AlertPanel from "../components/alert-panel";

export default function DashboardPage() {
  return (
    <main>
      <Header />

      <div>
        <StatCard title="Nodes Online" value="0" />
        <StatCard title="DNS Queries" value="0" />
        <StatCard title="Latency" value="0ms" />
        <StatCard title="Health" value="100%" />
      </div>

      <NodeTable />
      <HealthChart />
      <AlertPanel />
    </main>
  );
}
