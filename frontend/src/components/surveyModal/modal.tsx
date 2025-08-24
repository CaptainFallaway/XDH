import { useSignal } from '@preact/signals';
import type { ModalProps } from './types';
import { api } from '@wails/go/models';

import FileInput from './fileInput';

export default function Modal({ onCancel, onSubmit, modalId }: ModalProps) {
  const selectedFile = useSignal<string | null>(null);

  selectedFile.subscribe((value) => console.log('Selected file', value));

  const handleSubmit = async (e: Event) => {
    if (!selectedFile.value) {
      e.preventDefault();
      console.error('No file selected');
      return;
    }

    const form = e.target as HTMLFormElement;

    console.log('Form submitted', form);

    const dto: api.Survey = {
      surveyor: form.surveyor.value,
      date: Date.now() * 1000,
      location: form.location.value,
      instrumentSerial: form.instrumentSerial.value,
      westCoastFlag: form.westCoastFlag.value === 'true',
    };

    await onSubmit(dto, selectedFile.value!);
    selectedFile.value = null;
  };

  const handleCancel = async (e: Event) => {
    e.preventDefault();
    const modal = document.getElementById(modalId) as HTMLDialogElement;
    modal?.close();
    onCancel();
  };

  return (
    <dialog id={modalId} className="modal">
      <form onSubmit={handleSubmit} method="dialog" className="flex modal-box flex-col space-x-2">
        <div className="flex flex-col space-y-3">
          <fieldset className="fieldset bg-base-200 border-base-300 rounded-box w-full border p-4">
            <legend className="fieldset-legend">Mät Schema</legend>

            <label className="label">Mätförättare</label>
            <input type="text" className="input" name="surveyor" placeholder="John Doe" />

            <label className="label">Plats</label>
            <input type="text" className="input" name="location" placeholder="Lomma" />

            <label className="label">XRF Instrument - Serienummer</label>
            <input type="text" className="input" name="instrumentSerial" />

            <label className="label">Kust</label>
            <select name="westCoastFlag" className="select">
              <option value="true">Västkusten</option>
              <option value="false">Östkusten</option>
            </select>
          </fieldset>
          <FileInput selectedFile={selectedFile} />
        </div>
        <div className="modal-action flex w-full">
          <button type="submit" className="btn">
            Kör
          </button>
          <button type="button" className="btn" onClick={handleCancel}>
            Stäng
          </button>
        </div>
      </form>
    </dialog>
  );
}
