import React from 'react'
import {
    Layout, Menu, Breadcrumb, Icon,List,Skeleton,Button,Row, Col,Input
  } from 'antd';

  import {
    BrowserRouter as Router,
    Route,
    Link,
    Switch,
    NavLink,
    Redirect
  } from 'react-router-dom'

  import moment from 'moment';


  const {
    Header, Content, Footer, Sider,
  } = Layout;
  const SubMenu = Menu.SubMenu;
  const list = [{
    id:'3',  
    name:'test1'
  },{
    id:'4',  
    name:'test2'
  },{
    id:'5',
    name:'test1'
},{
    id:'6',
    name:'test2'
},{
    id:'7',
    name:'test1'
},{
    id:'8',
    name:'test2'
},{
    id:'9',
    name:'test1'
},{
    id:'10',
    name:'test2'
},{
    id:'11',
    name:'test1'
},{
    id:'12',
    name:'test2'
}]
  class TestLogDetailPage extends React.Component {
    state = {
    };
    render() {
      return (
        <Content>
            <br></br>
            <Header style={{ background: '#fff' }} >
                {/* <Row>
                    <Col xs={12} sm={12} md={18}></Col>
                </Row> */}
                <h4 style={{position:'relative'}}>
                    <span>任务运行详情页</span>
                    {`${this.props.match.params.id}`}
                    {/* <Button onClick={(e)=>{this.props.history.go(-1)}} style={{position:'absolute',right:0,top:'15px'}} icon="left">返回</Button> */}
                </h4>
            </Header>
            <div style={{ padding: 24, background: '#fff', minHeight: 360 }}>
                <div>
                    <span>任务名称：</span>
                    <span>name</span>
                </div>
                <div>
                    <span>运行时间：</span>
                    <span>{moment().format('LLLL')}</span>
                </div>
                <div>
                    <div>
                        <span>运行结果：</span>
                        <Button size='small'>重新运行</Button>
                    </div>
                    <div>
                        <Input.TextArea rows={6}></Input.TextArea>
                    </div>

                </div>
            </div>
        </Content>)
    }
}
export default TestLogDetailPage;
