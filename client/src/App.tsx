import { Toaster } from "sonner";
import "./App.css";
import TodoList from "./components/Todo";


function App() {
  return (
    <main>
      <TodoList />
      <Toaster />
    </main>

  )
}

export default App
