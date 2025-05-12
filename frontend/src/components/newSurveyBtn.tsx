import type { ComponentChildren } from 'preact';
import type { api } from '../lib/wailsjs/go/models';

const modalId = 'new-survey-modal';

export type onCancel = () => void | Promise<void>;
export type onSubmit = (dto: api.Survey, path: string) => void | Promise<void>;

export interface NewSurveyBtnProps {
  onCancel: onCancel;
  onSubmit: onSubmit;
  children: ComponentChildren;
  [key: string]: any;
}

export interface ModalProps {
  onCancel: onCancel;
  onSubmit: onSubmit;
}

function Modal({ onCancel, onSubmit }: ModalProps) {
  const handleSubmit = async (e: Event) => {
    const form = e.target as HTMLFormElement;

    console.log('Form submitted', form);

    const dto: api.Survey = {
      surveyor: form.surveyor.value,
      date: Date.now() * 1000,
      location: form.location.value,
      instrumentSerial: form.instrumentSerial.value,
    };

    onSubmit(dto);
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
        <div className="flex flex-col">
          <label htmlFor="surveyor">Surveyor</label>
          <input
            type="text"
            id="surveyor"
            name="surveyor"
            className="input input-bordered w-full max-w-xs"
            placeholder="John Flow"
          />
          <label htmlFor="location">Location</label>
          <input
            type="text"
            id="location"
            name="location"
            className="input input-bordered w-full max-w-xs"
            placeholder="Lomma"
          />
          <label htmlFor="instrumentSerial">Instrument Serial</label>
          <input
            type="text"
            id="instrumentSerial"
            name="instrumentSerial"
            className="input input-bordered w-full max-w-xs"
            placeholder="123456789"
          />
        </div>
        <div className="modal-action">
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

function NewSurveyBtn({ onCancel, onSubmit, children, ...props }: NewSurveyBtnProps) {
  const openModal = () => {
    const modal = document.getElementById(modalId) as HTMLDialogElement;
    modal?.showModal();
  };

  return (
    <>
      <Modal onCancel={onCancel} onSubmit={onSubmit} />
      <button onClick={openModal} {...props}>
        {children}
      </button>
    </>
  );
}

export default NewSurveyBtn;
