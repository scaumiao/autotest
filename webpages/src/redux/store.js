import { createStore } from 'redux'
import taskApp from './reducer/taskReducer'
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
} from './action/task.js'
let store = createStore(taskApp)
// 打印初始状态
console.log(store.getState())

// 每次 state 更新时，打印日志
// 注意 subscribe() 返回一个函数用来注销监听器
const unsubscribe = store.subscribe(() =>
  console.log(store.getState())
)

// 发起一系列 action
store.dispatch(addTask({id:'1',task:'addtask1'}))
store.dispatch(addTask({id:'2',task:'addtask2'}))
store.dispatch(addTask({id:'3',task:'addtask3'}))
store.dispatch(modifyTask({id:'2',task:'addtask3_modify'}))
store.dispatch(deleteTask({id:'2'}))
store.dispatch(clearTasks())
store.dispatch(setTasks([{id:'233333'}]))

// 停止监听 state 更新
unsubscribe();

export default store;