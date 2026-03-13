# Azure Web Server Architecture Diagram

```mermaid
graph TB
    subgraph Internet["🌐 Internet"]
        Users[Users/Clients]
    end

    subgraph Azure["☁️ Azure Cloud - westus3"]
        subgraph RG["📦 Resource Group<br/>${labelPrefix}-A05-RG"]
            
            PublicIP["🌍 Public IP Address<br/>${labelPrefix}A05PublicIP<br/>Dynamic Allocation"]
            
            subgraph VNet["🔷 Virtual Network<br/>${labelPrefix}A05Vnet<br/>10.0.0.0/16"]
                
                subgraph Subnet["📡 Subnet<br/>${labelPrefix}A05Subnet<br/>10.0.1.0/24"]
                    
                    VM["🖥️ Linux Virtual Machine<br/>${labelPrefix}A05VM<br/>Ubuntu 22.04 LTS<br/>Standard_B1s<br/>Apache2 Web Server"]
                    
                    NIC["🔌 Network Interface<br/>${labelPrefix}A05Nic<br/>Private IP: Dynamic"]
                    
                end
            end
            
            NSG["🛡️ Network Security Group<br/>${labelPrefix}A05SG<br/>Rules:<br/>✓ SSH (Port 22)<br/>✓ HTTP (Port 80)"]
            
        end
    end

    Users -->|HTTP/SSH| PublicIP
    PublicIP -->|Connected to| NIC
    NIC -->|Attached to| VM
    NSG -->|Protects| NIC
    VM -.->|Installed via cloud-init| Apache[Apache2]

    style Internet fill:#e1f5ff
    style Azure fill:#f0f0f0
    style RG fill:#fff4e6
    style VNet fill:#e6f3ff
    style Subnet fill:#f0f8ff
    style VM fill:#d4edda
    style NSG fill:#fff3cd
    style PublicIP fill:#cce5ff
```

## Architecture Components

### Network Layer
- **Virtual Network (VNet)**: 10.0.0.0/16 CIDR block
- **Subnet**: 10.0.1.0/24 CIDR block within the VNet
- **Public IP**: Dynamic allocation for external access

### Compute Layer
- **Virtual Machine**: Standard_B1s size running Ubuntu 22.04 LTS
- **Web Server**: Apache2 installed automatically via cloud-init script

### Security Layer
- **Network Security Group**: Controls inbound traffic
  - SSH access on port 22
  - HTTP access on port 80
- **Network Interface**: Connects VM to the network with security group attached

### Resource Organization
- **Resource Group**: Logical container for all resources in westus3 region

## Traffic Flow

1. **Inbound Traffic**: Internet → Public IP → Network Interface → Virtual Machine
2. **Security**: Network Security Group filters traffic at the NIC level
3. **Web Service**: Apache2 serves HTTP requests on port 80
4. **Management**: SSH access available on port 22 for administration
