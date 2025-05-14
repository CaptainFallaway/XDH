import * as app from '@wails/go/app/App';
import { Signal } from '@preact/signals';
import { Upload } from 'lucide-preact';

export default function FileInput({ selectedFile }: { selectedFile: Signal<string | null> }) {
  const handleClick = async (e: Event) => {
    e.preventDefault();
    selectedFile.value = await app.OpenFileDialog();
  };

  const extractFileName = (path: string) => {
    if (path.includes('/')) {
      return path.split('/').pop() || '';
    } else if (path.includes('\\')) {
      return path.split('\\').pop() || '';
    }
  };

  return (
    <button
      onClick={handleClick}
      className="flex flex-col items-center border-4 border-base-300 border-dashed rounded-md py-6 space-y-3"
    >
      <Upload size={48} />
      <p>{selectedFile.value ? extractFileName(selectedFile.value) : 'Select File'}</p>
    </button>
  );
}
