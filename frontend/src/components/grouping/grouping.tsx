import { models } from '@wails/go/models';
import { useSignal } from '@preact/signals';
import { ChevronDown, ChevronUp } from 'lucide-preact';
import { formatLongDate } from '@utils/dates';

import ScansTable from './scanTable';
import MetalSummary from './metalStat';
import Notes from './notes';

export interface GroupingProps {
  grouping: models.Grouping;
  metal: string;
  westCoastFlag: boolean;
  [key: string]: any;
}

function Expanded({ grouping, westCoastFlag }: { grouping: models.Grouping; westCoastFlag: boolean }) {
  return (
    <div className="mt-4" onClick={(e) => e.stopPropagation()}>
      <MetalSummary violations={grouping.violations} westCoastFlag={westCoastFlag} />
      <Notes notes={grouping.errorNotes} />
      <ScansTable scans={grouping.validScans} unit={grouping.unit} westCoastFlag={westCoastFlag} />
    </div>
  );
}

export function SmallInfo({ title, text }: { title: string; text: string }) {
  return (
    <div class="justify-start flex flex-col space-y-2 p-2 whitespace-nowrap">
      <strong>{title}</strong>
      <p>{text}</p>
    </div>
  );
}

export default function Grouping({ grouping, metal, westCoastFlag, ...props }: GroupingProps) {
  let border = 'border-success';

  if (grouping.violations[metal] === 1) {
    border = 'border-warning';
  } else if (grouping.violations[metal] > 1) {
    border = 'border-error';
  }

  const expanded = useSignal(false);

  return (
    <button
      onClick={(e) => {
        e.stopPropagation();
        expanded.value = !expanded.value;
      }}
      className={`${border} flex flex-col w-full p-4 bg-base-200 border rounded-md`}
      {...props}
    >
      <div className="flex w-full">
        <SmallInfo title="Båt Id" text={grouping.boatID} />
        <SmallInfo title="Mätförättare" text={grouping.operators.join(', ')} />
        <SmallInfo title="Första Mätning" text={formatLongDate(grouping.firstDate)} />
        <SmallInfo title="Sista Mätning" text={formatLongDate(grouping.lastDate)} />
        <div className="flex w-full items-center justify-end">
          {expanded.value ? <ChevronDown size={48} /> : <ChevronUp size={48} />}
        </div>
      </div>
      {expanded.value ? <Expanded grouping={grouping} westCoastFlag={westCoastFlag} /> : null}
    </button>
  );
}
