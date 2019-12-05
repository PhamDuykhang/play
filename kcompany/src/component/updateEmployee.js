import React, { Component } from "react";
import {
  Form,
  Input,
  Icon,
  Cascader,
  Select,
  Button,
  Tag,
  Divider
} from "antd";

import "./employeeFrom.css";
const { Option } = Select;
const ButtonGroup = Button.Group;

//get from API
const delivery = [
  {
    value: "dg1",
    label: "DG-4",
    children: [
      {
        value: "dc14",
        label: "DC-14",
        children: [
          {
            value: "kbtgkplus",
            label: "KBTG-K+ Shop"
          },
          {
            value: "vts",
            label: "Vital Suite"
          },
          {
            value: "vtq",
            label: "Vital QIP"
          },
          {
            value: "tsi",
            label: "Transcend Insight"
          }
        ]
      },
      {
        value: "dg4",
        label: "DG-4",
        children: [
          {
            value: "acatel",
            label: "Ancatel Router Device"
          },
          {
            value: "humanad",
            label: "Humana Dev"
          },
          {
            value: "osx",
            label: "Oxigen Solution"
          }
        ]
      }
    ]
  }
];

class EmployeeFrom extends Component {
  constructor(props) {
    super(props);
    this.state = {
      isLocked: true,
      skills:[],
      employee: {
        name: "La Ngoc Nguyen",
        department: "DC14",
        tech_skill: [
          { k: "java", value: "Java" },
          { k: "golang", value: "Golang" }
        ],
        id: "K11-2233",
        addr: {
          street: "Hang Gon",
          Province: "Can Tho"
        },
        department:["dg1", "dc14", "kbtgkplus"]
      },
      //will be got from API
      systemSkill: [
        { k: "java", value: "Java" },
        { k: "golang", value: "Golang" },
        { k: "gcc", value: "Google Cloud" },
        { k: "java", value: "Java" },
        { k: "ivy", value: "Ivy" },
        { k: "test", value: "Tester" }
      ]
    };
  }

  componentDidMount() {
    this.state.employee.tech_skill.map(ski => this.state.skills.push(ski.k));
  }

  handleSubmit = e => {
    e.preventDefault();
    this.props.form.validateFieldsAndScroll((err, values) => {
      if (!err) {
        console.log("Received values of form: ", values);
      }
    });
  };
  unLockForm = () => {
    this.setState({
      isLocked: false
    });
  };
  handleConfirmBlur = e => {
    const { value } = e.target;
    this.setState({ confirmDirty: this.state.confirmDirty || !!value });
  };

  handleAreaClick(e, label, option) {
    e.stopPropagation();
  }
  //Used later
  displayRenderAddr = (labels, selectedOptions) =>
    labels.map((label, i) => {
      console.log(selectedOptions);
      const option = selectedOptions[i];
      if (i === labels.length - 1) {
        return (
          <span key={option.value}>
            {label}
            <a onClick={e => this.handleAreaClick(e, label, option)}></a>
          </span>
        );
      }
      return <span key={option.value}>{label} - </span>;
    });
  render() {
    const { getFieldDecorator } = this.props.form;
    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 8 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 10 }
      }
    };
    const tailFormItemLayout = {
      wrapperCol: {
        xs: {
          span: 24,
          offset: 0
        },
        sm: {
          span: 16,
          offset: 8
        }
      }
    };
    const listDropDown = this.state.systemSkill.map(skill => (
      <Option key={skill.k}>
        <Tag color="green">
          <Icon type="tag" />
          <Divider type="vertical"></Divider>
          {skill.value}
        </Tag>
      </Option>
    ));
    const prefixSelector = getFieldDecorator("prefix", {
      initialValue: "(+84)Việt Nam"
    })(
      <Select disabled={this.state.isLocked} mod="tags" style={{ width: 120 }}>
        <Option value="84">(+84)Việt Nam</Option>
        <Option value="66">(+66)Thái Lan</Option>
      </Select>
    );
    return (
      <Form {...formItemLayout} onSubmit={this.handleSubmit}>
        <Form.Item label="Employee ID">
          {getFieldDecorator("employeeID", {
            initialValue: this.state.employee.id
          })(<Input disabled={true} />)}
        </Form.Item>
        <Form.Item label="Employee Name" hasFeedback>
          {getFieldDecorator("employeeName", {
            initialValue: this.state.employee.name
          })(<Input disabled={this.state.isLocked} />)}
        </Form.Item>
        <Form.Item label="Technical Skill">
          {getFieldDecorator("skill", {
            initialValue: this.state.skill
          })(
            <Select
              mode="multiple"
              style={{ width: '100%' }}
              allowClear={true}
              disabled={this.state.isLocked}
            >
              {listDropDown}
            </Select>
          )}
        </Form.Item>
        <Form.Item label="Department">
          {getFieldDecorator("department", {
            initialValue: ["dg1", "dc14", "kbtgkplus"]
          })(
            <Cascader
              disabled={this.state.isLocked}
              changeOnSelect={true}
              options={delivery}
            />
          )}
        </Form.Item>
        <Form.Item label="Phone Number">
          {getFieldDecorator("phone", {
            rules: [
              { required: true, message: "Please input your phone number!" }
            ],
            initialValue: this.state.employee.phone
          })(
            <Input
              disabled={this.state.isLocked}
              addonBefore={prefixSelector}
              style={{ width: "100%" }}
            />
          )}
        </Form.Item>
        <Form.Item label="Employee Address">
          {getFieldDecorator("addr", {
            initialValue:
              this.state.employee.addr.street +
              ", " +
              this.state.employee.addr.Province
          })(<Input disabled={this.state.isLocked} />)}
        </Form.Item>
        <Form.Item {...tailFormItemLayout}>
          <ButtonGroup>
            <Button
              disabled={this.state.isLocked}
              type="primary"
              htmlType="submit"
            >
              Update
            </Button>
            <Button onClick={this.unLockForm} icon="edit">
              Edit
            </Button>
          </ButtonGroup>
        </Form.Item>
      </Form>
    );
  }
}
const UpdateEmployee = Form.create({ name: "updating-employee-from" })(EmployeeFrom);
export default UpdateEmployee;
