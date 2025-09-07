# Task Owner
---
## Dogan Akyurek
## akyurek.dogan.dgn@gmail.com
---
# Home24 | Website Analyzer

## 📦 Requirements

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
| Core architecture implementation & edge cases   | 1h       |
| Service development                             | 2h       |
| Testing & documentation                         | 1h       |
| Deployment & hand-off                           | 0.5h     |

---

### ✅ Completed Tasks

- [x] Requirement gathering & edge case analysis
- [x] Project layout & git repository initialization
- [x] System design and architecture skeleton
- [x] MVP & dry run.
- [x] Implementation improvements
- [x] UI Implementation
- [ ] Unit test coverage check & increase
- [x] Project documentation
- [x] Project delivery to Home24 team

---

### ⚠️ Known Limitations / Missing parts

- Logging is implemented using uber/zap. Due to time constraints, I initialized it with the production configuration using the Sugared Logger. Most logs could be converted to DEBUG level for more detailed output.
- Different functionalities are implemented as plugins, which can be activated on demand via configuration (see main.go, lines 54–57).
- Asynchronous processing is demonstrated in a basic way to show concurrency is possible. With more time, the system could be made more efficient (see html_analyzer.go, lines 25–40).
- Unit, E2E, and performance tests were not added due to time constraints.
- Concurrency could be improved further; several steps, such as link processing and validation, can be optimized for better efficiency.
- Test cases were primarily performed on https://bbc.com since it is a feature-rich site. Larger websites were not pushed to avoid excessive processing time.
- Containerization is skipped.
- CI/CD pipeline is skipped.

---

## 🚀 Getting Started

### 🧰 Prerequisites

- Golang must be installed on local system
- You can check the task.pdf in the source code for task details.

---

## 🧪 Running the Project

### Step 1 – Setup

- Golang must be installed in the system.

---

### Step 2 – Run Options

- Please use the make file command to run the project `make run`
- Following steps will be done automatically when you use the make file however better to check
    - Please copy the `.env.dist file as .env` on your local.
    - Please copy the `.config.yml.dist as config.yml` on your local
- The web page must be opened automatically otherwise please visit the `localhost:8080` if it's not due to your browser or local security settings please visit manually.
---

### Step 3 – Test the Application

- You can use following websites to check the application (simpler to extensive)
    - example.com -> Very clean so better to check simple steps.
    - bbc.com -> I used for local tests, pretty packed website.

## 📸 Screenrecordings

https://github.com/user-attachments/assets/f1c1b7b0-441c-4684-944e-59da70f6fdfe

https://github.com/user-attachments/assets/55685665-46fb-4ed2-81da-026b2cd7987e

https://github.com/user-attachments/assets/0aea6257-9891-4dd1-985f-76c6b7e44c82

## 🤝 Contributing

This is a sample project submitted to Home24.

---

## 📬 Contact
---

*Dogan Akyurek*
*akyurek.dogan.dgn@gmail.com*

For questions or feedback, feel free to open an issue or reach out via *akyurek.dogan.dgn@gmail.com*

---
