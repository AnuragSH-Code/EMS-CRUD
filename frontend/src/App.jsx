import Card from "./components/comman/Card";
import Profile from "./components/comman/Profile";
import SideBarButton from "./components/comman/SideBarButton";
import Team from "./components/Team";

function App() {
  return (
    <main className="grid grid-cols-1 lg:grid-cols-5 h-screen">
      
      <aside className="col-span-1 flex flex-col justify-between bg-white-off p-6 lg:flex lg:flex-col lg:justify-between">
        <header>
          <h1 className="text-3xl font-medium py-4">EMS-BREVO</h1>
          <p className="text-xl font-medium py-4 text-red-500">NOT WORKING AUTH</p>
          <nav className="flex flex-col gap-2">
            <SideBarButton text="Dashboard" />
            <SideBarButton text="Team" />
            <SideBarButton text="Connect" />
          </nav>
        </header>

        <footer className="flex flex-col gap-2">
          <Profile />
          <SideBarButton text="Settings" />
          <SideBarButton text="Logout" />
        </footer>
      </aside>


      <section className="col-span-1 lg:col-span-4 h-full overflow-scroll">
        <Team />
      </section>
    </main>
  );
}

export default App;
