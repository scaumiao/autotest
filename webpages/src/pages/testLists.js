import React from 'react'
import {
    Layout, Menu, Breadcrumb, Icon,List,Skeleton,Button,Row, Col
  } from 'antd';

  import {
    BrowserRouter as Router,
    Route,
    Link,
    Switch,
    NavLink,
    Redirect
  } from 'react-router-dom'
  const {
    Header, Content, Footer, Sider,
  } = Layout;
  const SubMenu = Menu.SubMenu;
  const list = [{
    id:'3',  
    name:'test1'
  },{
    id:'3',  
    name:'test2'
  },{
    id:'3',
    name:'test1'
},{
    id:'3',
    name:'test2'
},{
    id:'3',
    name:'test1'
},{
    id:'3',
    name:'test2'
},{
    id:'3',
    name:'test1'
},{
    id:'3',
    name:'test2'
},{
    id:'3',
    name:'test1'
},{
    id:'3',
    name:'test2'
}]
  class TestListsPage extends React.Component {
    state = {
    };
    render() {
      return (
          <Content>
            <Header style={{ background: '#fff' }} >
                {/* <Row>
                    <Col xs={12} sm={12} md={18}></Col>
                </Row> */}
                <h3 style={{position:'relative'}}>
                    <span>任务列表</span>
                    <NavLink to={`${this.props.match.path}/createTest`}>
                        <Button style={{position:'absolute',right:0,top:'15px'}} icon="plus">新建任务</Button>
                    </NavLink>
                </h3>
            </Header>
            <Content style={{ margin: '0 16px' }}>
            <Breadcrumb style={{ margin: '16px 0' }}>
                <Breadcrumb.Item>任务</Breadcrumb.Item>
                <Breadcrumb.Item>任务列表</Breadcrumb.Item>
            </Breadcrumb>
            <div style={{ padding: 24, background: '#fff', minHeight: 360 }}>
                <List
                    className="demo-loadmore-list"
                    itemLayout="horizontal"
                    dataSource={list}
                    pagination={{
                        onChange: (page) => {
                        //   console.log(page);
                        },
                        pageSize: 5,
                      }}
                    renderItem={item => (
                        <List.Item actions={
                            [
                                <NavLink to={`${this.props.match.path}/editTest/`+item.id}>
                                    <Button size='small'>修改</Button>
                                </NavLink>,
                                <NavLink to={`${this.props.match.path}/startTest/`+item.id}>
                                    <Button size='small'>启动</Button>
                                </NavLink>,
                                <NavLink to={`${this.props.match.path}/delTest/`+item.id}>
                                    <Button size='small'>删除</Button>
                                </NavLink>,
                                <NavLink to={`${this.props.match.path}/logTest/`+item.id}>
                                    <Button size='small'>运行记录</Button>
                                </NavLink>
                            ]
                        }>
                            <List.Item.Meta
                                title={<a href="https://ant.design">{item.name}</a>}
                            />
                        </List.Item>
                    )}
                />
            </div>
            </Content>
          </Content>)
        }
}
export default TestListsPage;
