import React from 'react';
import { LayoutDashboard, Users, CreditCard, TrendingUp, Package } from 'lucide-react';

function Sidebar({ activeTab, setActiveTab }) {
  return (
    <aside className="w-full md:w-64 glass border-b md:border-b-0 md:border-r border-white/10 p-4 md:p-6 flex flex-col gap-4 md:gap-8 z-20">
      <div className="flex items-center gap-3 px-2">
        <div className="w-10 h-10 bg-primary rounded-xl flex items-center justify-center shadow-lg shadow-primary/20">
          <TrendingUp size={24} className="text-white" />
        </div>
        <span className="text-xl font-bold tracking-tight text-white">BankSave</span>
      </div>

      <nav className="flex flex-row md:flex-col gap-2 overflow-x-auto pb-2 md:pb-0 scrollbar-hide">
        <NavItem 
          active={activeTab === 'dashboard'} 
          onClick={() => setActiveTab('dashboard')} 
          icon={<LayoutDashboard size={20}/>} 
          label="Dashboard" 
        />
        <NavItem 
          active={activeTab === 'customers'} 
          onClick={() => setActiveTab('customers')} 
          icon={<Users size={20}/>} 
          label="Customers" 
        />
        <NavItem 
          active={activeTab === 'accounts'} 
          onClick={() => setActiveTab('accounts')} 
          icon={<CreditCard size={20}/>} 
          label="Accounts" 
        />
        <NavItem 
          active={activeTab === 'packages'} 
          onClick={() => setActiveTab('packages')} 
          icon={<Package size={20}/>} 
          label="Packages" 
        />
      </nav>
    </aside>
  );
}

function NavItem({ icon, label, active, onClick }) {
  return (
    <button 
      onClick={onClick}
      className={`flex items-center gap-2 md:gap-3 px-3 py-2 md:px-4 md:py-3 rounded-xl transition-all whitespace-nowrap ${
        active ? 'bg-primary text-white shadow-lg shadow-primary/20' : 'hover:bg-white/5 text-secondary hover:text-white'
      }`}
    >
      {icon}
      <span className="font-medium text-sm md:text-base">{label}</span>
    </button>
  );
}

export default Sidebar;
