import Modal from './modal';
import type { NewSurveyBtnProps } from './types';

const modalId = 'new-survey-modal';

export default function NewSurveyBtn({ onCancel, onSubmit, children, ...props }: NewSurveyBtnProps) {
  // This is probably not the best way to do this, but it works for now
  const openModal = () => {
    const modal = document.getElementById(modalId) as HTMLDialogElement;
    modal?.showModal();
  };

  return (
    <>
      <Modal onCancel={onCancel} onSubmit={onSubmit} modalId={modalId} />
      <button onClick={openModal} {...props}>
        {children}
      </button>
    </>
  );
}
