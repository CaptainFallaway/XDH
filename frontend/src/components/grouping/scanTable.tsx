import { formatLongDate } from '@utils/dates';
import { models } from '@wails/go/models';

export default function ScansTable({
  scans,
  unit,
  westCoastFlag,
}: {
  scans: models.Scan[];
  unit: string;
  westCoastFlag: boolean;
}) {
  return (
    <div class="overflow-x-auto">
      <table class="table bg-base-100 table-zebra w-full">
        <thead>
          <tr>
            <th>Actions</th>
            <th>Reading</th>
            <th>Duration</th>
            <th>Operator</th>
            <th>Date/Time</th>
            {westCoastFlag ? (
              <>
                <th>Pb ({unit})</th>
                <th>Sn ({unit})</th>
              </>
            ) : (
              <>
                <th>Cu ({unit})</th>
                <th>Zn ({unit})</th>
              </>
            )}
            <th>Enabled</th>
          </tr>
        </thead>
        <tbody>
          {scans.map((scan, i) => (
            <tr
              key={scan.reading}
              // className={!scan.enabled ? "opacity-50" : ""}
            >
              <td className="flex gap-1">
                <button
                  className="btn btn-xs"
                  // onClick={() => moveScanUp(i)}
                  // disabled={i === 0}
                >
                  ↑
                </button>
                <button
                  className="btn btn-xs"
                  // onClick={() => moveScanDown(i)}
                  disabled={i === scans.length - 1}
                >
                  ↓
                </button>
              </td>
              <td>{scan.reading}</td>
              <td>{scan.duration.toFixed(2)}</td>
              <td>{scan.operator}</td>
              <td>
                {formatLongDate(scan.date)}
                {/* <br />
                <span className="text-xs opacity-70">{formatLongDate(scan.date)}</span> */}
              </td>
              {westCoastFlag ? (
                <>
                  <td className={scan.violations['pb'] ? 'text-error font-bold' : ''}>{scan.pb.toFixed(2)}</td>
                  <td className={scan.violations['sn'] ? 'text-error font-bold' : ''}>{scan.sn.toFixed(2)}</td>
                </>
              ) : (
                <>
                  <td className={scan.violations['cu'] ? 'text-error font-bold' : ''}>{scan.cu.toFixed(2)}</td>
                  <td className={scan.violations['zn'] ? 'text-error font-bold' : ''}>{scan.zn.toFixed(2)}</td>
                </>
              )}
              <td>
                <input
                  type="checkbox"
                  className="toggle toggle-primary toggle-sm"
                  checked={true}
                  // onChange={() => toggleScan(i)}
                />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
