import React from 'react'
import {
  BrowserRouter as Router,
  Route,
  Link,
  Switch,
  Redirect
} from 'react-router-dom'
import { LocaleProvider } from 'antd';
import zh_CN from 'antd/lib/locale-provider/zh_CN';
import moment from 'moment';
import 'moment/locale/zh-cn';
import IndexPage from './pages/index'
moment.locale('zh-cn');

const App = () => (
  <LocaleProvider locale={zh_CN}>
    <Router>
      <Switch>
        <Route path="/lists" component={IndexPage}/>
        <Route path="/about" render={() => <div>about</div>}/>
        <Redirect to='/lists'/>
        <Route component={IndexPage}/>
      </Switch>
    </Router>
  </LocaleProvider>
)
export default App