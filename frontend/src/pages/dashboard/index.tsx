import { Col, Row, Card, Statistic, Space, Table, Tag } from "antd";
import {
  AuditOutlined,
  UserOutlined,
  PieChartOutlined,
  StockOutlined,
} from "@ant-design/icons";
import type { ColumnsType } from "antd/es/table";

interface DataType {
  key: string;
  name: string;
  age: number;
  address: string;
  tags: string[];
}

const columns: ColumnsType<DataType> = [
  {
    title: "ลำดับ",
    dataIndex: "ID",
    key: "id",
  },
  {
    title: "ชื่อ",
    dataIndex: "FirstName",
    key: "firstname",
  },
  {
    title: "นามสกุล",
    dataIndex: "LastName",
    key: "lastname",
  },
  {
    title: "อีเมล",
    dataIndex: "Email",
    key: "email",
  },
  {
    title: "เบอร์โทร",
    dataIndex: "Phone",
    key: "phone",
  },
];

const data: DataType[] = [];

export default function index() {
  return (
    <>
      <Row gutter={[16, 16]}>
        <Col xs={24} sm={24} md={24} lg={24} xl={24}>
          <h2>แดชบอร์ด</h2>
        </Col>
        <Col xs={24} sm={24} md={24} lg={24} xl={24}>
          <Card style={{ backgroundColor: "#F5F5F5" }}>
            <Row gutter={[16, 16]}>
              <Col xs={24} sm={24} md={12} lg={12} xl={6}>
                <Card
                  bordered={false}
                  style={{
                    boxShadow: "rgba(100, 100, 111, 0.2) 0px 7px 29px 0px",
                  }}
                >
                  <Statistic
                    title="จำนวน"
                    value={1800}
                    prefix={<StockOutlined />}
                  />
                </Card>
              </Col>
              <Col xs={24} sm={24} md={12} lg={12} xl={6}>
                <Card
                  bordered={false}
                  style={{
                    boxShadow: "rgba(100, 100, 111, 0.2) 0px 7px 29px 0px",
                  }}
                >
                  <Statistic
                    title="จำนวน"
                    value={200}
                    valueStyle={{ color: "black" }}
                    prefix={<AuditOutlined />}
                  />
                </Card>
              </Col>
              <Col xs={24} sm={24} md={12} lg={12} xl={6}>
                <Card
                  bordered={false}
                  style={{
                    boxShadow: "rgba(100, 100, 111, 0.2) 0px 7px 29px 0px",
                  }}
                >
                  <Statistic
                    title="จำนวน"
                    value={3000}
                    valueStyle={{ color: "black" }}
                    prefix={<PieChartOutlined />}
                  />
                </Card>
              </Col>
              <Col xs={24} sm={24} md={12} lg={12} xl={6}>
                <Card
                  bordered={false}
                  style={{
                    boxShadow: "rgba(100, 100, 111, 0.2) 0px 7px 29px 0px",
                  }}
                >
                  <Statistic
                    title="จำนวน"
                    value={10}
                    valueStyle={{ color: "black" }}
                    prefix={<UserOutlined />}
                  />
                </Card>
              </Col>
            </Row>
          </Card>
        </Col>
        <Col xs={24} sm={24} md={24} lg={24} xl={24}>
          <h3>ผู้ใช้งานล่าสุด</h3>
        </Col>
        <Col xs={24} sm={24} md={24} lg={24} xl={24}>
          <Table columns={columns} dataSource={data} />
        </Col>
      </Row>
    </>
  );
}
