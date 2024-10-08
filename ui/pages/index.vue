<template>
  <div class="todo-main">
    <h1>TODOリスト</h1>
    <div v-if="statusMessage" class="status-message">{{ statusMessage }}</div>
    <div class="input-group">
      <input v-model="newTask" placeholder="Enter new task" @keyup.enter="addTodo">
      <div class="input-dropdown">
        <select v-model="selectedPriority" >
          <option disabled value="">Select a task</option>
          <option v-for="priority in priorities" :key="priority.id" :value="priority.value">
            {{ priority.value }}
          </option>
        </select>
      </div>
      <button @click="addTodo">Addition</button>
    </div>
    <div v-if="todos.length > 0">
      <div v-for="todo in todos" :key="todo.ID" class="todo-item">
        <input
            v-if="todo.isEditing" v-model="todo.Task" class="edit-input" @blur="editTodo(todo)"
            @keyup.enter="editTodo(todo)">
        <span  v-else :class="{ 'done-task': todo.Status === 'done' }" @click="enableEdit(todo)">{{
            todo.Task
          }}</span>
        <div class="buttons">
          <button :class="{ 'done': todo.Status === 'done' }" @click="updateStatus(todo)">
            ✔️
          </button>
          <button class="delete-button" @click="deleteTodo(todo.ID)">🗑️</button>
        </div>
      </div>
    </div>
    <div v-else>
      <p>There are no tasks.</p>
    </div>


    <h3>Done List</h3>
<!--    done list-->
    <div v-if="doneTodos.length > 0">
      <div v-for="doneTodo in doneTodos" :key="doneTodo.ID" class="todo-item">
        <input
            v-if="doneTodo.isEditing" v-model="doneTodo.Task" class="edit-input" @blur="editTodo(doneTodo)"
            @keyup.enter="editTodo(doneTodo)">
        <span v-else  @click="enableEdit(doneTodo)">{{
            doneTodo.Task
          }}</span>
      </div>
    </div>
    <div v-else>
      <p>There are no completed tasks.</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newTask: '',
      todos: [],
      doneTodos: [],
      statusMessage: '',
      selectedPriority: '',
      priorities: [
        { id: 1, value: 'Normal' },
        { id: 2, value: 'Medium' },
        { id: 3, value: 'High' },
      ],
    };
  },
  mounted() {
    this.fetchTodos();
    this.fetchDoneTodos()
  },
  methods: {
    async fetchTodos() {
      try {
        const response = await fetch(`/api/v1/todos?sorted=true`, {});
        if (!response.ok) throw new Error(`Failed to get todo list. statusCode: ${response.status}`);
        response.json().then(data => {
          this.todos = data.data;
        });
      } catch (error) {
        console.error(error);
        this.statusMessage = 'Failed to get task';
      }
    },
    async fetchDoneTodos() {
      try {
        const response = await fetch(`/api/v1/todos?status=done`, {});
        if (!response.ok) throw new Error(`Failed to get todo list. statusCode: ${response.status}`);
        response.json().then(data => {
          this.doneTodos = data.data;
          console.log("adasdad", this.todos)
        });
      } catch (error) {
        console.error(error);
        this.statusMessage = 'Failed to get task';
      }
    },
    async addTodo() {
      if (this.newTask.trim() === '') return;
      if (this.selectedPriority.trim() === '') return;
      try {
        const response = await fetch('/api/v1/todos', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            task: this.newTask,
            priority: this.selectedPriority,
            Status: 'created'
          })
        });

        if (!response.ok) throw new Error(`Failed to create todo. statusCode: ${response.status}`);

        response.json().then(data => {
          this.todos.push(data.data);
          this.newTask = '';
          this.selectedPriority = '';
          this.statusMessage = 'task added';
        });
      } catch (error) {
        console.error('Error creating todo:', error);
        this.statusMessage = 'Failed to create task';
      }
    },
    enableEdit(todo) {
      todo.isEditing = true;
    },
    async editTodo(todo) {
      todo.isEditing = false;

      try {
        const response = await fetch(`/api/v1/todos/${todo.ID}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            task: todo.Task
          })
        });

        if (!response.ok) throw new Error(`Failed to edit todo. statusCode: ${response.status}`);

        this.statusMessage = 'Task edited';
      } catch (error) {
        console.error('Error editing todo:', error);
        this.statusMessage = 'Failed to edit task';
      }
    },
    async updateStatus(todo) {
      try {
        const response = await fetch(`/api/v1/todos/${todo.ID}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            Status: todo.Status === 'done' ? 'created' : 'done'
          })
        });

        if (!response.ok) throw new Error(`Failed to update todo Status. statusCode: ${response.status}`);

        todo.Status = todo.Status === 'done' ? 'created' : 'done';
        this.statusMessage = 'Task status changed';
        await this.fetchDoneTodos();
      } catch (error) {
        console.error('Error updating todo Status:', error);
        this.statusMessage = 'Failed to update status';
      }
    },
    async deleteTodo(id) {
      try {
        const response = await fetch(`/api/v1/todos/${id}`, {
          method: 'DELETE',
          headers: {
            'Content-Type': 'application/json',
          },
        });

        if (!response.ok) throw new Error(`Failed to update todo Status. statusCode: ${response.status}`);

        this.todos = this.todos.filter(todo => todo.ID !== id);
        this.statusMessage = 'Task deleted';
        await this.fetchDoneTodos()
      } catch (error) {
        console.error('Error deleting todo:', error);
        this.statusMessage = 'Failed to delete task';
      }
    }
  }
};
</script>

<style scoped>
.todo-main {
  max-width: 400px;
  margin: 20px auto;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  background-color: #fff;
}

.input-group {
  display: flex;
  margin-bottom: 20px;
}

input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  margin-right: 10px;
}

button {
  padding: 10px;
  border: none;
  background-color: #333;
  color: #fff;
  border-radius: 4px;
  cursor: pointer;
}

.todo-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.buttons button {
  background-color: #f0f0f0;
  color: #333;
  margin-left: 5px;
  border-radius: 4px;
  padding: 5px 10px;
  transition: background-color 0.3s, color 0.3s;
}

.buttons button.done {
  background-color: #333;
  color: #fff;
}

.buttons button.done::before {
  color: #fff;
}

.buttons button.delete-button {
  color: white;
}

.status-message {
  margin-bottom: 20px;
  padding: 10px;
  background-color: #f0f0f0;
  border-radius: 4px;
}

.done-task {
  text-decoration: line-through;
}

.edit-input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.input-dropdown {
  display: flex;
  padding-right: 5px;
  border-radius: 8px;
}
</style>