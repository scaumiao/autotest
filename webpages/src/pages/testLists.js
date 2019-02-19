import React from 'react'
import {
    Layout, Menu, Breadcrumb, Icon,List,Skeleton,
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
      name:'test1'
  },{
      name:'test2'
  },{
    name:'test1'
},{
    name:'test2'
},{
    name:'test1'
},{
    name:'test2'
},{
    name:'test1'
},{
    name:'test2'
},{
    name:'test1'
},{
    name:'test2'
}]
  class TestListsPage extends React.Component {
    state = {
    };
    render() {
      return (
          <Content>
            <Header style={{ background: '#fff' }} >
                <h2>测试列表</h2>
            </Header>
            <Content style={{ margin: '0 16px' }}>
            <Breadcrumb style={{ margin: '16px 0' }}>
                <Breadcrumb.Item>User</Breadcrumb.Item>
                <Breadcrumb.Item>Bill</Breadcrumb.Item>
            </Breadcrumb>
            <div style={{ padding: 24, background: '#fff', minHeight: 360 }}>
                <List
                    className="demo-loadmore-list"
                    itemLayout="horizontal"
                    dataSource={list}
                    pagination={{
                        onChange: (page) => {
                          console.log(page);
                        },
                        pageSize: 3,
                      }}
                    renderItem={item => (
                    <List.Item actions={[<a>edit</a>, <a>more</a>]}>
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
