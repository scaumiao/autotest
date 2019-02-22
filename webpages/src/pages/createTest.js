import React from 'react'
import {
    Layout, Menu, Breadcrumb, Icon,List,Skeleton,
    Button,Row, Col,Input
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
  const { TextArea } = Input;

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
  class CreateTestPage extends React.Component {
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
                    <span>新建任务</span>
                    <Button onClick={(e)=>{this.props.history.go(-1)}} style={{position:'absolute',right:0,top:'15px'}} icon="left">返回</Button>
                </h3>
            </Header>
            <Content style={{ margin: '0 16px' }}>
            <Breadcrumb style={{ margin: '16px 0' }}>
                <Breadcrumb.Item>测试任务</Breadcrumb.Item>
                <Breadcrumb.Item>新建任务</Breadcrumb.Item>
            </Breadcrumb>
                <div style={{ padding: 24, background: '#fff', minHeight: 360 }}>
                    <Row>
                        <Col span={6}>
                            <label style={{lineHeight:'32px',textAlign:'right'}} for='taskName'>任务名:</label>
                        </Col>
                        <Col span={18}>
                            <Input id='taskName' placeholder="请输入任务名" />
                        </Col>
                    </Row>
                    <Row style={{marginTop:'15px'}}>
                        <Col span={24}>
                            <label style={{lineHeight:'32px',textAlign:'right'}} for='taskCode'>脚本代码:</label>
                            <Button size='small' style={{marginLeft:'15px'}}>运行</Button>
                        </Col>
                        <Col span={24} style={{marginTop:'7px'}}>
                            <TextArea id='taskCode' rows={8} />
                        </Col>
                    </Row>
                    <Row style={{marginTop:'15px'}}>
                        <Col span={6}>
                            <label style={{lineHeight:'32px',textAlign:'right'}} for='runPeriod'>运行周期</label>
                        </Col>
                        <Col span={18}>
                            <Input id='runPeriod' placeholder="请输入运行周期" />
                        </Col>
                    </Row>
                    <Row style={{marginTop:'15px'}}>
                    <Col span={3} offset={9}>
                        <Button
                        onClick={e => {
                            this.props.history.go(-1);
                        }}
                        >
                        取消新增
                        </Button>
                    </Col>
                    <Col span={3}>
                        <Button type="primary">确认新增</Button>
                    </Col>
                    </Row>
                </div>
            </Content>
          </Content>)
        }
}
export default CreateTestPage;
