import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import CampaignsPage from './CampaignsPage';
import CreateCampaignPage from './CreateCampaignPage';
import './App.css';

export default function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<CampaignsPage />} />
        <Route path="/create" element={<CreateCampaignPage />} />
      </Routes>
    </Router>
  );
}