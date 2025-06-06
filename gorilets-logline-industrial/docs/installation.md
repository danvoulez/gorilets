# Installation

## Prerequisites
- Python 3.8+
- Node.js 16+

## Steps
1. Install Python dependencies:
   ```bash
   pip install lark sqlalchemy websocket-client
   ```
2. Install Node dependencies:
   ```bash
   cd src/frontend && npm install
   ```
3. Initialize the database:
   ```bash
   cd src/logline
   python -c "print('init_db placeholder')"
   ```
4. Start the frontend:
   ```bash
   npm run dev
   ```
