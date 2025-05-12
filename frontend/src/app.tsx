import { signal } from '@preact/signals';

import * as app from './lib/wailsjs/go/app/App';
import { api, models } from './lib/wailsjs/go/models';
import NewSurveyBtn from './components/newSurveyBtn';

const selectedSurvey = signal<models.Survey | null>(null);
const surveys = signal<models.Survey[]>([]);
const groupings = signal<models.Grouping[]>([]);

selectedSurvey.subscribe(async (value) => {
  if (!value) {
    groupings.value = [];
    return;
  }

  const temp = await app.GetGroupings(value.groupingIds);
  groupings.value = temp;
});

async function createSurvey(dto: api.Survey) {
  console.log('Creating survey', dto);
  return;

  const path = await app.OpenFileDialog();
  if (!path) {
    console.error('No path selected');
    return;
  }

  selectedSurvey.value = await app.CreateSurvey(dto, path);
}

async function getSurveys() {
  const temp = await app.GetSurveys();
  console.log('Suveys', temp);
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
            New Survey
          </NewSurveyBtn>
          <button className="btn" onClick={getSurveys}>
            Get Surveys
          </button>
          <button className="btn" onClick={deleteSurvey}>
            Delete Current Survey
          </button>
        </div>
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
      </div>

      {groupings.value.map((grouping) => (
        <p>{grouping.boatID}</p>
      ))}
    </>
  );
}
