import { 
    ADD_TASK,
    DELETE_TASK,
    MODIFY_TASK,
    SET_TASKS,
    CLEAR_TASKS,
    addTask,
    deleteTask,
    modifyTask ,
    setTasks,
    clearTasks
} from '../action/task.js'

const initialState = {
  tasks: []
};

function taskApp(state = initialState, action) {
  // 这里暂不处理任何 action，
  // 仅返回传入的 state。
  switch (action.type) {
    case ADD_TASK:
      return Object.assign({}, state, {
        tasks: [...state.tasks,action.task]
      })
    case DELETE_TASK:
      {
        let _tasks = []
        state.tasks.map((task,index)=>{
            if (task.id!==action.task.id) {
                _tasks.push(task)
            }
        })

        return Object.assign({},state,{
            tasks:_tasks
        })
      }
    case MODIFY_TASK:
      {
        let _tasks_ = [];
        _tasks_ = state.tasks.map((task,index)=>{
            if (task.id ==action.task.id) {
                return action.task
            }else{
                return task
            }
        })
        return Object.assign({},state,{
            tasks:_tasks_
        })
      }
    case SET_TASKS:
        return Object.assign({},state,{
            tasks:action.tasks
        })
    case CLEAR_TASKS:
        return Object.assign({},state,{
            tasks:[]
        })
    default:
      return state
  }
}

export default taskApp;