import { AlertTriangle } from 'lucide-preact';

export default function Notes({ notes }: { notes: string[] }) {
  if (!notes || notes.length === 0) {
    return null;
  }

  return (
    <div class="alert alert-warning mb-4">
      <AlertTriangle size={18} />
      <div>
        <h3 class="font-bold">Notes</h3>
        <ul class="list-disc list-inside">
          {notes.map((note, index) => (
            <li key={index}>{note}</li>
          ))}
        </ul>
      </div>
    </div>
  );
}
