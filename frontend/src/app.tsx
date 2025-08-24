import { signal } from '@preact/signals';

import * as app from '@wails/go/app/App';
import { api, models } from './lib/wailsjs/go/models';
import NewSurveyBtn from '@components/surveyModal/newSurveryBtn';
import Grouping from '@components/grouping/grouping';

const selectedSurvey = signal<models.Survey | null>(null);
const surveys = signal<models.Survey[]>([]);
const groupings = signal<models.Grouping[]>([]);
const sortingMetal = signal<string>('Pb');

selectedSurvey.subscribe(async (value) => {
  if (!value) {
    groupings.value = [];
    return;
  }

  console.log('Selected survey', value);

  const temp = await app.GetGroupings(value.groupingIds, sortingMetal.value);

  console.log('Groupings', temp);

  groupings.value = temp;
});

async function createSurvey(dto: api.Survey, path: string) {
  console.log('Creating survey', dto);

  selectedSurvey.value = await app.CreateSurvey(dto, path);
}

async function getSurveys() {
  const temp = await app.GetSurveys();
  console.log('Surveys', temp);
  surveys.value = temp;
}

async function deleteSurvey() {
  if (!selectedSurvey.value) {
    console.error('No survey selected');
    return;
  }

  await app.DeleteSurvey(selectedSurvey.value.uid);
  selectedSurvey.value = null;
}

export function App() {
  return (
    <>
      <div className="flex content-center justify-center">
        <div className="flex flex-col space-y-1">
          <NewSurveyBtn className="btn" onCancel={() => console.log('Cancelled')} onSubmit={createSurvey}>
            Nytt Mätschema
          </NewSurveyBtn>
          <button className="btn" onClick={getSurveys}>
            Ladda mätscheman
          </button>
          <button className="btn" onClick={deleteSurvey}>
            Ta bort nuvarande mätschema
          </button>
        </div>
        <div className="flex flex-col">
          <select
            className="select"
            onChange={async (e) => {
              const survey = surveys.value.find((s) => s.uid === e.currentTarget.value);
              if (survey) selectedSurvey.value = survey;
            }}
          >
            <option selected>Pick a survey</option>
            {surveys.value.map((survey) => (
              <option key={survey.uid} value={survey.uid}>
                {survey.surveyor} - {survey.location} - {survey.date * 1000} - {survey.instrumentSerial}
              </option>
            ))}
          </select>
          {/* {selectedSurvey.value?.westCoastFlag ? (
            <select name="" id=""></select>
          ) : (
            <select name="" id=""></select>
          )
            } */}
        </div>
      </div>

      <div className={'flex flex-col space-y-2 p-4'}>
        {groupings.value.map((grouping) => (
          <Grouping
            key={grouping.boatID}
            grouping={grouping}
            metal={sortingMetal.value}
            westCoastFlag={selectedSurvey.value?.westCoastFlag || false}
          />
        ))}
      </div>
    </>
  );
}
