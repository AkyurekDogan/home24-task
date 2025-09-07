# Task Owner
---
## Dogan Akyurek
## akyurek.dogan.dgn@gmail.com
---
# Home24 | Website Analyzer

Gets the following inputs from the customer
- url
    - check validity of the url if not fail fast.
    - check the url accesibility if not fail fast.

Returns the following details
- response
    - What HTML version has the document?
    - What is the page title?
    - How many headings of what level are in the document?
    - How many internal and external links are in the document? Are there any inaccessible links and how many?
    - Does the page contain a login form?

---

## 📦 Project Layout

### ⏱ Milestones (avg hours)

| Task                                            | Duration |
|-------------------------------------------------|----------|
| Requirement analysis & architecture design      | 1h       |
| Core architecture implementation & edge cases   | 2h       |
| Service development                             | 5h       |
| Testing & documentation                         | 0h       |
| Terraform & Kubernetes integration              | 0.5h     |
| Deployment & hand-off                           | 0.5h     |

---

### ✅ Completed Tasks

- [x] Requirement gathering & edge case analysis
- [x] Project layout & git repository initialization
- [ ] System design and architecture skeleton
- [ ] MVP & dry run.
- [ ] Implementation improvements
- [ ] UI Implementation
- [ ] Unit test coverage check & increase
- [ ] Project documentation
- [ ] Project delivery to Home24 team

---

### ⚠️ Known Limitations / Missing parts

- I used the uber/zap for logging but due to time constraints I initiated as sugar version with production config so I did not used DEBUG mode most of the logs can be converted to debug mode.
- I made the different functionalities as plugins so we can activate these plugins by need using config. (file main.go, line 54:57 )
- Async system can be more compdlext and efficient but due to time constraint I tried to show something that this operations can be doable in concurrent way. (file html_analyzer.go, line 25:40 )

---

## 🚀 Getting Started

### 🧰 Prerequisites

TBD

---

## 🧪 Running the Project

### Step 1 – Setup

---

### Step 2 – Run Options

#### ✅ Option 1: Developer Mode (Recommended)

---

#### 🐳 Option 2: Docker Compose
---

---

### Step 3 – Test the API


#### 🧪 Option 1: grpcurl (CLI)
```

```

#### 🧪 Option 2: Postman

You can also use [Postman](https://www.postman.com/) 

---

## 🧱 Architecture Diagram

---

## 📸 Screenrecordings


## 🤝 Contributing

This is a sample project submitted to Home24.

---

## 📬 Contact
---

*Dogan Akyurek*
*akyurek.dogan.dgn@gmail.com*

For questions or feedback, feel free to open an issue or reach out via *akyurek.dogan.dgn@gmail.com*

---
