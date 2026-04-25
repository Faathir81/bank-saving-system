import React, { useState, useEffect } from 'react';
import Sidebar from './components/Sidebar';
import DashboardView from './views/DashboardView';
import CustomerView from './views/CustomerView';
import AccountView from './views/AccountView';
import { dataService } from './services/dataService';
import { Plus } from 'lucide-react';

function App() {
  const [activeTab, setActiveTab] = useState('dashboard');
  const [data, setData] = useState({
    customers: [],
    accounts: [],
    depositoTypes: []
  });
  const [loading, setLoading] = useState(true);

  const loadAllData = async () => {
    setLoading(true);
    try {
      const result = await dataService.getDashboardData();
      setData(result);
    } catch (err) {
      console.error("Failed to fetch data", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadAllData();
  }, []);

  return (
    <div className="flex min-h-screen bg-background text-slate-200">
      <Sidebar activeTab={activeTab} setActiveTab={setActiveTab} />

      <main className="flex-1 p-10 overflow-y-auto">
        <header className="mb-10 flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold capitalize tracking-tight text-white">{activeTab}</h1>
            <p className="text-secondary mt-1 text-sm">Managing Bank Saving System Core.</p>
          </div>
          <div className="flex gap-4">
             <div className="px-4 py-2 bg-white/5 rounded-lg border border-white/10 text-xs font-mono text-secondary">
                Server: <span className="text-emerald-400">Online</span>
             </div>
          </div>
        </header>

        {loading ? (
          <div className="flex items-center justify-center h-64">
            <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
          </div>
        ) : (
          <>
            {activeTab === 'dashboard' && (
              <DashboardView customers={data.customers} accounts={data.accounts} />
            )}
            {activeTab === 'customers' && (
              <CustomerView customers={data.customers} onRefresh={loadAllData} />
            )}
            {activeTab === 'accounts' && (
              <AccountView 
                accounts={data.accounts} 
                customers={data.customers} 
                depositoTypes={data.depositoTypes} 
                onRefresh={loadAllData} 
              />
            )}
          </>
        )}
      </main>
    </div>
  );
}

export default App;
