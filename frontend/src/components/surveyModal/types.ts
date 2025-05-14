import type { ComponentChildren } from 'preact';
import type { api } from '@wails/go/models';

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
  modalId: string;
}