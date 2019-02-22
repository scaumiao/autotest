export const ADD_TASK = 'ADD_TASK';
export const DELETE_TASK = 'DELETE_TASK';
export const MODIFY_TASK = 'MODIFY_TASK';
export const SET_TASKS = 'SET_TASKS';
export const CLEAR_TASKS = 'CLEAR_TASKS';


export function addTask(task) {
    return {
      type: ADD_TASK,
      task
    }
}
  
export function deleteTask(task) {
    return {
      type: DELETE_TASK,
      task
    }
}

export function modifyTask(task) {
    return {
      type: MODIFY_TASK,
      task
    }
}
export function setTasks(tasks) {
    return {
      type: SET_TASKS,
      tasks
    }
}
export function clearTasks() {
    return {
      type: CLEAR_TASKS,
    }
}