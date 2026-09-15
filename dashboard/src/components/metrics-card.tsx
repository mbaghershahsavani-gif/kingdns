export default function MetricsCard({
  title,
  value,
}: {
  title: string;
  value: string;
}) {
  return (
    <div>
      <span>{title}</span>
      <strong>{value}</strong>
    </div>
  );
}
