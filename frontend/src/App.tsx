import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Layout from './Layout';
import CampaignsPage from './CampaignsPage';
import CreateCampaignPage from './CreateCampaignPage';
import MyCampaignsPage from './MyCampaignsPage';
import MyDonationsPage from './MyDonationsPage';
import './App.css';

export default function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<CampaignsPage />} />
          <Route path="/create" element={<CreateCampaignPage />} />
          <Route path="/my-campaigns" element={<MyCampaignsPage />} />
          <Route path="/my-donations" element={<MyDonationsPage />} />
        </Routes>
      </Layout>
    </Router>
  );
}