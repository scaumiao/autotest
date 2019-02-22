import React from 'react'
import {
    Layout, Menu, Breadcrumb, Icon,
  } from 'antd';
import {
BrowserRouter as Router,
Route,
Link,
Switch,
NavLink,
Redirect
} from 'react-router-dom'
import TestListsPage from './testLists'
import CreateTestPage from './createTest'
import LogTestPage from './logTest'
//测试store
import store from '../redux/store.js'
const {
Header, Content, Footer, Sider,
} = Layout;

const SubMenu = Menu.SubMenu;

class IndexPage extends React.Component {
    state = {
        collapsed: false,
    };

    onCollapse = (collapsed) => {
        console.log(collapsed);
        this.setState({ collapsed });
    }

    render() {
        // console.log(this.props.match)
        return (
        <Layout style={{ minHeight: '100vh' }}>
            <Sider
            collapsible
            collapsed={this.state.collapsed}
            onCollapse={this.onCollapse}
            >
            <div className="logo" />
            <Menu theme="dark" defaultSelectedKeys={['lists']} defaultOpenKeys={['test']} mode="inline">
                {/* <Menu.Item key="2">
                    <NavLink to="/about">
                        <Icon type="pie-chart" />
                        <span>测试列表</span>
                    </NavLink>
                </Menu.Item> */}
                <SubMenu
                key="test"
                title={<span><Icon type="team" /><span>测试</span></span>}
                >
                <Menu.Item key="lists">
                    <NavLink to="/lists">
                        <Icon type="pie-chart" />
                        <span>测试列表</span>
                    </NavLink>
                </Menu.Item>
                <Menu.Item key="about">
                    <NavLink to={`${this.props.match.path}/about`}>
                        <Icon type="pie-chart" />
                        <span>about</span>
                    </NavLink>
                </Menu.Item>
                </SubMenu>
            </Menu>
            </Sider>
            <Layout>
            <Switch>
                <Route exact path={`${this.props.match.path}`} component={TestListsPage}/>
                <Route path={`${this.props.match.path}/about`} render={() => <Content>about:{`${this.props.match.path}/about`}</Content>}/>
                <Route path={`${this.props.match.path}/createTest`} component={CreateTestPage} />
                <Route path={`${this.props.match.path}/editTest/:id`} component={CreateTestPage} />
                <Route path={`${this.props.match.path}/logTest/:id`} component={LogTestPage} />
                <Redirect to={`${this.props.match.path}`}/>
                <Route component={TestListsPage}/>
            </Switch>
            {/* <Header style={{ background: '#fff', padding: 0 }} />
            <Content style={{ margin: '0 16px' }}>
                <Breadcrumb style={{ margin: '16px 0' }}>
                <Breadcrumb.Item>User</Breadcrumb.Item>
                <Breadcrumb.Item>Bill</Breadcrumb.Item>
                </Breadcrumb>
                <div style={{ padding: 24, background: '#fff', minHeight: 360 }}>
                Bill is a cat.
                </div>
            </Content> */}
            <Footer style={{ textAlign: 'center' }}>
                Ant Design ©2018 Created by Ant UED
            </Footer>
            </Layout>
        </Layout>
        );
    }
}
export default IndexPage;