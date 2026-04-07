import { useState, useMemo, useEffect } from "react";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { Card } from "@/components/ui/card";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";


import { Search, Trash2, Pencil, Eye, Plus, ClipboardList } from "lucide-react";
import { create, getAll } from "@/api/todos";
import { toast } from "sonner";

interface Todo {
  id: string;
  title: string;
  description: string;
  completed: boolean;
}

const Index = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  // const [isFetchingTodos, setIsFetchingTodos] = useState(false);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [search, setSearch] = useState("");
  const [viewTodo, setViewTodo] = useState<Todo | null>(null);
  const [editTodo, setEditTodo] = useState<Todo | null>(null);
  const [editTitle, setEditTitle] = useState("");
  const [editDescription, setEditDescription] = useState("");

  // useEffect(() => {
  //   getAll().then((res)=> {
  //     const 
  //     setTodos
  //   }).catch((err) => {
  //     console.error("Failed to fetch todos:", err);
  //   });
  // }, [setTodos]);

  const filteredTodos = useMemo(() => {
    const q = search.toLowerCase();
    if (!q) return todos;
    return todos.filter(
      (t) =>
        t.title.toLowerCase().includes(q) ||
        t.description.toLowerCase().includes(q)
    );
  }, [todos, search]);

  const addTodo = async () => {
    if (!title.trim()) return;

    await create(title, description).then((res) => {
      console.log(res.json())
      toast.success("Created")
    }).catch((err) => {
      console.log(err)
      toast.error(err.message ?? "FAILED TO CREATE")
    })

    // setTodos((prev) => [
    // {
    // id: crypto.randomUUID(),
    // title: title.trim(),
    // description: description.trim(),
    // completed: false,
    // },
    // ...prev,
    // ]);
    // setTitle("");
    // setDescription("");
  };

  const toggleTodo = (id: string) => {
    setTodos((prev) =>
      prev.map((t) => (t.id === id ? { ...t, completed: !t.completed } : t))
    );
  };

  const deleteTodo = (id: string) => {
    setTodos((prev) => prev.filter((t) => t.id !== id));
  };

  const deleteAll = () => setTodos([]);

  const openEdit = (todo: Todo) => {
    setEditTodo(todo);
    setEditTitle(todo.title);
    setEditDescription(todo.description);
  };

  const saveEdit = () => {
    if (!editTodo || !editTitle.trim()) return;
    setTodos((prev) =>
      prev.map((t) =>
        t.id === editTodo.id
          ? { ...t, title: editTitle.trim(), description: editDescription.trim() }
          : t
      )
    );
    setEditTodo(null);
  };

  return (
    <div className="min-h-screen bg-background py-8 px-4">
      {/* Header */}
      <div className="flex items-center justify-center gap-3 w-full mb-5">
        <ClipboardList className="h-8 w-8 text-primary" />
        <h1 className="text-3xl font-bold text-foreground">Todo App</h1>
      </div>

      <div className="mx-auto space-y-6 md:grid grid-cols-8 gap-2">
        {/* Add Form */}
        <div className="col-span-3">
          <Card className="p-4 space-y-3 sticky! top-2">
            <Input
              placeholder="Title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              onKeyDown={(e) => e.key === "Enter" && addTodo()}
            />
            <Textarea
              placeholder="Description (optional)"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              className="min-h-[60px]"
            />
            <Button onClick={addTodo} disabled={!title.trim()} className="w-full gap-2">
              <Plus className="h-4 w-4" /> Add Todo
            </Button>
          </Card>
        </div>


        {/* Search & Bulk Actions */}
        <div className="space-y-2 col-span-5">

          <div className="flex gap-2 sticky top-0 bg-background/80 backdrop-blur-sm z-10 px-1 py-2">
            <div className="relative flex-1">
              <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
              <Input
                placeholder="Search todos..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                className="pl-9"
              />
            </div>
            {todos.length > 0 && (
              <AlertDialog>
                <AlertDialogTrigger asChild>
                  <Button variant="destructive" size="sm" className="gap-1 shrink-0">
                    <Trash2 className="h-4 w-4" /> Delete All
                  </Button>
                </AlertDialogTrigger>
                <AlertDialogContent>
                  <AlertDialogHeader>
                    <AlertDialogTitle>Delete all todos?</AlertDialogTitle>
                    <AlertDialogDescription>
                      This will permanently remove all {todos.length} todo(s). This action cannot be undone.
                    </AlertDialogDescription>
                  </AlertDialogHeader>
                  <AlertDialogFooter>
                    <AlertDialogCancel>Cancel</AlertDialogCancel>
                    <AlertDialogAction onClick={deleteAll}>Delete All</AlertDialogAction>
                  </AlertDialogFooter>
                </AlertDialogContent>
              </AlertDialog>
            )}
          </div>

          {/* Todo List */}
          <div className="space-y-2">
            {filteredTodos.length === 0 && (
              <div className="py-12 text-center text-muted-foreground">
                {todos.length === 0
                  ? "No todos yet. Add one above!"
                  : "No todos match your search."}
              </div>
            )}
            {filteredTodos.map((todo) => (
              <Card
                key={todo.id}
                className={`p-4 flex items-start gap-3 transition-opacity ${todo.completed ? "opacity-60" : ""
                  }`}
              >
                <Checkbox
                  checked={todo.completed}
                  onCheckedChange={() => toggleTodo(todo.id)}
                  className="mt-1"
                />
                <div className="flex-1 min-w-0">
                  <p
                    className={`font-medium text-foreground ${todo.completed ? "line-through text-muted-foreground" : ""
                      }`}
                  >
                    {todo.title}
                  </p>
                  {todo.description && (
                    <p className="text-sm text-muted-foreground truncate">
                      {todo.description}
                    </p>
                  )}
                </div>
                <div className="flex gap-1 shrink-0">
                  <Button
                    variant="ghost"
                    size="icon"
                    className="h-8 w-8"
                    onClick={() => setViewTodo(todo)}
                  >
                    <Eye className="h-4 w-4" />
                  </Button>
                  <Button
                    variant="ghost"
                    size="icon"
                    className="h-8 w-8"
                    onClick={() => openEdit(todo)}
                  >
                    <Pencil className="h-4 w-4" />
                  </Button>
                  <AlertDialog>
                    <AlertDialogTrigger asChild>
                      <Button variant="ghost" size="icon" className="h-8 w-8 text-destructive">
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    </AlertDialogTrigger>
                    <AlertDialogContent>
                      <AlertDialogHeader>
                        <AlertDialogTitle>Delete this todo?</AlertDialogTitle>
                        <AlertDialogDescription>
                          "{todo.title}" will be permanently removed.
                        </AlertDialogDescription>
                      </AlertDialogHeader>
                      <AlertDialogFooter>
                        <AlertDialogCancel>Cancel</AlertDialogCancel>
                        <AlertDialogAction onClick={() => deleteTodo(todo.id)}>
                          Delete
                        </AlertDialogAction>
                      </AlertDialogFooter>
                    </AlertDialogContent>
                  </AlertDialog>
                </div>
              </Card>
            ))}
          </div>
        </div>
      </div>

      {/* View Dialog */}
      <Dialog open={!!viewTodo} onOpenChange={() => setViewTodo(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{viewTodo?.title}</DialogTitle>
          </DialogHeader>
          <p className="text-sm text-muted-foreground whitespace-pre-wrap">
            {viewTodo?.description || "No description."}
          </p>
          <p className="text-xs text-muted-foreground">
            Status: {viewTodo?.completed ? "Completed ✓" : "Pending"}
          </p>
        </DialogContent>
      </Dialog>

      {/* Edit Dialog */}
      <Dialog open={!!editTodo} onOpenChange={() => setEditTodo(null)}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Edit Todo</DialogTitle>
          </DialogHeader>
          <div className="space-y-3">
            <Input
              value={editTitle}
              onChange={(e) => setEditTitle(e.target.value)}
              placeholder="Title"
            />
            <Textarea
              value={editDescription}
              onChange={(e) => setEditDescription(e.target.value)}
              placeholder="Description"
            />
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setEditTodo(null)}>
              Cancel
            </Button>
            <Button onClick={saveEdit} disabled={!editTitle.trim()}>
              Save
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default Index;
