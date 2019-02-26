import React from 'react'
import {
    Layout, Menu, Breadcrumb, Icon,List,Skeleton,
    Button,Row, Col,Input,notification
  } from 'antd';

  import {
    BrowserRouter as Router,
    Route,
    Link,
    Switch,
    NavLink,
    Redirect
  } from 'react-router-dom'
  import axios from 'axios';
  const {
    Header, Content, Footer, Sider,
  } = Layout;
  const SubMenu = Menu.SubMenu;
  const { TextArea } = Input;
  class CreateTestPage extends React.Component {
    state = {
        task:{
        },
        type:'新增',
        test:''
    };
    _input_={

    };
    handleSure(e){
        let {history} = this.props
        if (this.props.match.params.id) {
            axios.put('/v1/task/'+this.props.match.params.id, this.state.task)
          .then((res)=>{
            if(res.status=='200'){
                console.log(res);
                notification.success({
                    message:this.state.type+'任务成功！'
                })
                history.push({pathName:'/lists/editTest/'+res.data.id})
            }
          })
          .catch(function (error) {
            console.log(error);
            notification.error({
                message:this.state.type+'任务失败！'
            })
          });
        }else{
            axios.post('/v1/task', this.state.task)
            .then((res)=>{
              if(res.status=='200'){
                  console.log(res);
                  notification.success({
                      message:this.state.type+'任务成功！'
                  })
                  history.push({pathname:'/lists/editTest/'+res.data.id})
                //   history.go('-1');
              }
            })
            .catch(function (error) {
              console.log(error);
              notification.error({
                  message:this.state.type+'任务失败！'
              })
            });
        }
    }
    
    handleTaskNameChange(e){
        // if(e && e.target && e.target.value){
            let value = e.target.value;
            console.log(value);
            // if(e && e.target && e.target.value){
                // let value = e.target.value;
                this.setState(function(state,prop){
                    let task = Object.assign({},{...state.task},{name:value})
                    console.log(task)
                    return {
                        task
                    }
                })
            // }
        // }
    }
    handleTaskScriptChange(e){
        // if(e && e.target && e.target.value){
            let value = e.target.value;
            this.setState(function(state,prop){
                let task = Object.assign({},{...state.task},{script:value})
                return {
                    task
                }
            })
        // }
    }
    handleTaskPeriodChange(e){
        if(e && e.target && e.target.value){
            let value = e.target.value;
            this.setState({task:Object.assign({},this.state.task,{period:value})})
        }
    }
    // handleClick(e){
    //     let value = e.target.value;
    //     console.log(value);
    //     this.setState({
    //         test:value
    //     })

    // }
    getTask(taskId){
        console.log(this.props.match)
        axios.get('/v1/task/'+taskId)
        .then((res)=>{
            console.log(res)
            if(res.status=='200'){
                // notification.success({
                //     message:'获取任务成功！'
                // })
                this.setState({
                    task:res.data
                })
                // this._input_ = 
            }
        },(err)=>{
            console.error(err)
        })
        .catch((err)=>{
            notification.error({
                message:'获取任务失败！'
            })
        })
    }
    startTest(taskId){
        axios.post('/v1/task/'+taskId+'/run')
        .then((res)=>{
            console.log(res);

        })
        .catch((err)=>{
            notification.error({
                message:'启动失败！'
            })
        })
    }
    componentWillReceiveProps(nextProps){
        if (nextProps.match.params.id) {
            this.getTask(nextProps.match.params.id);
            this.setState({
                type:'修改'
            })
        }else{
            this.setState({
                type:'新增',
                task:{}
            })
        }
    }

    componentDidMount(){
        if (this.props.match.params.id) {
            this.getTask(this.props.match.params.id);
            this.setState({
                type:'修改'
            })
        }
    }
    render() {
        let {name,script} = this.state.task;
        let {test}=this.state;
      return (
          <Content>
            <Header style={{ background: '#fff' }} >
                {/* <Row>
                    <Col xs={12} sm={12} md={18}></Col>
                </Row> */}
                <h3 style={{position:'relative'}}>
                    <span>{this.state.type}任务</span>
                    <Button onClick={(e)=>{ this.props.history.push({pathname:'/lists'})}} style={{position:'absolute',right:0,top:'15px'}} icon="left">返回</Button>
                </h3>
            </Header>
            <Content style={{ margin: '0 16px' }}>
            <Breadcrumb style={{ margin: '16px 0' }}>
                <Breadcrumb.Item>测试任务</Breadcrumb.Item>
                <Breadcrumb.Item>{this.state.type}任务</Breadcrumb.Item>
            </Breadcrumb>
                <div style={{ padding: 24, background: '#fff', minHeight: 360 }}>
                    <Row>
                        <Col span={6}>
                            <label style={{lineHeight:'32px',textAlign:'right'}} >任务名:</label>
                        </Col>
                        <Col span={18}>
                            <Input id='taskName' placeholder="请输入任务名" value={this.state.task.name} allowClear onChange={this.handleTaskNameChange.bind(this)}/>
                        </Col>
                    </Row>
                    <Row style={{marginTop:'15px'}}>
                        {/* <Input value={test} onChange={this.handleClick.bind(this)}/> */}
                        <Col span={24}>
                            <label style={{lineHeight:'32px',textAlign:'right'}} for='taskScript'>脚本代码:</label>
                            <Button size='small' style={{marginLeft:'15px'}}>运行</Button>
                        </Col>
                        <Col span={24} style={{marginTop:'7px'}}>
                            <TextArea id='taskScript' rows={8} autosize={{minRows:8}} value={this.state.task.script} onChange={this.handleTaskScriptChange.bind(this)}/>
                        </Col>
                    </Row>
                    <Row style={{marginTop:'15px'}}>
                        <Col span={6}>
                            <label style={{lineHeight:'32px',textAlign:'right'}} for='runPeriod'>运行周期</label>
                        </Col>
                        <Col span={18}>
                            <Input id='runPeriod' placeholder="请输入运行周期"  value={this.state.task.period||''} onChange={this.handleTaskPeriodChange.bind(this)}/>
                        </Col>
                    </Row>
                    <Row style={{marginTop:'15px'}}>
                    <Col span={3} offset={9}>
                        <Button
                        onClick={e => {
                            this.props.history.go('-1');
                        }}
                        >
                        取消{this.state.type}
                        </Button>
                    </Col>
                    <Col span={3}>
                        <Button type="primary" onClick={this.handleSure.bind(this)}>确认{this.state.type}</Button>
                    </Col>
                    </Row>
                </div>
            </Content>
          </Content>)
        }
}
export default CreateTestPage;
