import React, { useState, useEffect } from 'react';
import { Bell, CheckCircle, AlertCircle } from 'lucide-react';
import Sidebar from './components/Sidebar';
import DashboardView from './views/DashboardView';
import CustomerView from './views/CustomerView';
import AccountView from './views/AccountView';
import DepositoView from './views/DepositoView';
import { dataService } from './services/dataService';

function App() {
  const [activeTab, setActiveTab] = useState('dashboard');
  const [customers, setCustomers] = useState([]);
  const [accounts, setAccounts] = useState([]);
  const [depositoTypes, setDepositoTypes] = useState([]);
  
  // Notification State
  const [toast, setToast] = useState(null);
  const [confirm, setConfirm] = useState(null);

  const fetchData = async () => {
    try {
      const [cRes, aRes, dRes] = await Promise.all([
        dataService.getCustomers(),
        dataService.getAccounts(),
        dataService.getDepositoTypes()
      ]);
      setCustomers(cRes.data || []);
      setAccounts(aRes.data || []);
      setDepositoTypes(dRes.data || []);
    } catch (err) {
      console.error("Fetch error:", err);
      showToast("Error fetching data from server", "error");
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  const showToast = (message, type = 'success') => {
    setToast({ message, type });
    setTimeout(() => setToast(null), 3000);
  };

  const handleConfirm = (message, onConfirm) => {
    setConfirm({ message, onConfirm });
  };

  const renderContent = () => {
    switch (activeTab) {
      case 'dashboard': return <DashboardView customers={customers} accounts={accounts} />;
      case 'customers': return <CustomerView customers={customers} onRefresh={fetchData} notify={showToast} confirm={handleConfirm} />;
      case 'accounts': return <AccountView accounts={accounts} customers={customers} depositoTypes={depositoTypes} onRefresh={fetchData} notify={showToast} confirm={handleConfirm} />;
      case 'packages': return <DepositoView depositoTypes={depositoTypes} onRefresh={fetchData} notify={showToast} confirm={handleConfirm} />;
      default: return <DashboardView customers={customers} accounts={accounts} />;
    }
  };

  return (
    <div className="flex min-h-screen bg-background text-slate-200 font-sans selection:bg-primary/30">
      <Sidebar activeTab={activeTab} setActiveTab={setActiveTab} />
      
      <main className="flex-1 p-8 overflow-y-auto">
        <header className="flex justify-between items-center mb-10">
          <div>
            <h1 className="text-3xl font-black text-white capitalize tracking-tight">{activeTab}</h1>
            <p className="text-secondary text-sm">Welcome back, Administrator</p>
          </div>
          <div className="flex gap-4">
             <button className="w-10 h-10 rounded-xl bg-white/5 border border-white/10 flex items-center justify-center text-secondary hover:text-primary transition-all relative">
                <Bell size={20} />
                <span className="absolute top-2 right-2 w-2 h-2 bg-accent rounded-full animate-pulse"></span>
             </button>
             <div className="w-10 h-10 rounded-xl bg-gradient-to-br from-primary to-blue-600 flex items-center justify-center font-bold text-white shadow-lg shadow-primary/20">
                A
             </div>
          </div>
        </header>

        {renderContent()}
      </main>

      {/* Toast Notification */}
      {toast && (
        <div className={`fixed bottom-8 right-8 z-[100] flex items-center gap-3 px-6 py-4 rounded-2xl border glass shadow-2xl animate-in slide-in-from-right fade-in duration-300 ${
          toast.type === 'error' ? 'border-red-500/50 bg-red-500/10' : 'border-emerald-500/50 bg-emerald-500/10'
        }`}>
          {toast.type === 'error' ? <AlertCircle className="text-red-400" /> : <CheckCircle className="text-emerald-400" />}
          <span className="font-medium text-white">{toast.message}</span>
        </div>
      )}

      {/* Confirmation Modal */}
      {confirm && (
        <div className="fixed inset-0 z-[100] flex items-center justify-center p-4">
          <div className="absolute inset-0 bg-background/80 backdrop-blur-sm" onClick={() => setConfirm(null)}></div>
          <div className="card w-full max-w-sm relative z-10 animate-in zoom-in-95 fade-in duration-200 shadow-2xl border border-white/10">
            <h3 className="text-xl font-bold mb-4 text-white">Confirm Action</h3>
            <p className="text-secondary mb-8">{confirm.message}</p>
            <div className="flex gap-3">
              <button className="flex-1 bg-white/5 hover:bg-white/10 p-3 rounded-xl transition-all" onClick={() => setConfirm(null)}>Cancel</button>
              <button className="flex-1 btn-primary" onClick={() => { confirm.onConfirm(); setConfirm(null); }}>Confirm</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default App;
