export default function StatCard({
  title,
  value,
}: {
  title: string;
  value: string;
}) {
  return (
    <div>
      <p>{title}</p>
      <strong>{value}</strong>
    </div>
  );
}
