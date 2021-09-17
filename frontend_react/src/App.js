// Import React dependencies.
import { SyncingEditor } from "./components/SyncingEditor";
import { Execute } from "./components/Execute";
import { ButtonCodeRun } from "./components/ButtonCodeRun";
const App = () => {
  localStorage.clear();
  return (
    <div>
      <ButtonCodeRun />
      <SyncingEditor />
    </div>
  );
};
export default App;
