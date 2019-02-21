import React from 'react'
import {
    Layout, Menu, Breadcrumb, Icon,List,Skeleton,Table, Divider, Tag,
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
  
  import moment from 'moment';

  import TestLogDetailPage from './testLogDetail'
  const {
    Header, Content, Footer, Sider,
  } = Layout;
  const SubMenu = Menu.SubMenu;
  const { TextArea } = Input;
  const { Column, ColumnGroup } = Table;

  class LogTestPage extends React.Component {
    state = {
        locations:[]
    };
    componentWillReceiveProps(nextProps) {
        // console.log('componentWillReceiveProps-------------------');
        // console.log(nextProps.location);
        // console.log(this.props.location);
        if (nextProps.location !== this.props.location) {
        // navigated!
            // this.setState({
            //     locations:[...this.state.locations,this.props.location]
            // })
            // console.log(this.state.locations)
        }
    }
    render() {
        const match = this.props.match;
        console.log(match)
        const columns = [{
            title: '时间',
            dataIndex: 'time',
            key: 'time',
            render: text => <a href="javascript:;">{text}</a>,
          }, {
            title: '历时',
            dataIndex: 'during',
            key: 'during',
          }, {
            title: '结果',
            dataIndex: 'result',
            key: 'result',
          }, {
            title: 'Tags',
            key: 'tags',
            dataIndex: 'tags',
            render: tags => (
              <span>
                {tags.map(tag => {
                  let color = tag.length > 5 ? 'geekblue' : 'green';
                  if (tag === 'loser') {
                    color = 'volcano';
                  }
                  return <Tag color={color} key={tag}>{tag.toUpperCase()}</Tag>;
                })}
              </span>
            ),
          }, {
            title: '操作',
            key: 'action',
            render: (text, record) => (
              <span>
                  <NavLink to={`${match.url}/detail/${record.id}`}>
                    查看详情
                  </NavLink>
                {/* <a href="javascript:;">查看详情 {record.name}</a> */}
                {/* <Divider type="vertical" />
                <a href="javascript:;">移除</a> */}
              </span>
            ),
          }];
          
          const data = [{
            id:'1',
            key: '1',
            time: moment().format('LLLL'),
            during: 32,
            result: true,
            tags: ['nice', 'developer'],
          }, {
            id:'2',
            key: '2',
            time: moment().format('LLLL'),
            during: 42,
            result: false,
            tags: ['loser'],
          }, {
            id:'3',
            key: '3',
            time: moment().format('LLLL'),
            during: 32,
            result: true,
            tags: ['cool', 'teacher'],
          }];
          const location = this.props.location;
        //   console.log(location)
        //   console.log(this.props.history)
        //   console.log(this.props.match)
      return (
          <Content>
            <Header style={{ background: '#fff' }} >
                {/* <Row>
                    <Col xs={12} sm={12} md={18}></Col>
                </Row> */}
                <h3 style={{position:'relative'}}>
                    <span>运行记录</span>
                    <Button onClick={(e)=>{this.props.history.push({pathname:'/lists'})}} style={{position:'absolute',right:0,top:'15px'}} icon="left">返回</Button>
                </h3>
            </Header>
            <Content style={{ margin: '0 16px' }}>
            <Breadcrumb style={{ margin: '16px 0' }}>
                <Breadcrumb.Item>测试任务</Breadcrumb.Item>
                <Breadcrumb.Item>运行记录</Breadcrumb.Item>
            </Breadcrumb>
                <div style={{ padding: 24, background: '#fff', minHeight: 360 }}>
                    <Table columns={columns} dataSource={data} />   
                </div>
                <div>
                        <Route exact path={`${this.props.match.path}/detail/:id`} component={TestLogDetailPage} />                    
                </div>
            </Content>
          </Content>)
        }
}
export default LogTestPage;
