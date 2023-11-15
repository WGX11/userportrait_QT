#include <QtCore>
#include <QSet>
#include "tool.h"
#include "collectorcontroller.h"

const int HeaderNum = 11;
static QStringList FileHeader{"事件类型", "事件时间","坐标",
                             "鼠标点击类型","鼠标点击按键","鼠标移动类型",
                             "键盘点击类型","键盘点击键",
                             "组件名","组件类型","组件附加信息"};


namespace qt_collector
{

class UserEventAnalyzer;
struct EventInfo;


class Agent
#ifndef Q_MOC_RUN
    final
#endif
    : public QObject
{
    Q_OBJECT
public:
    explicit Agent(QObject *parent = nullptr);
    ~Agent();
    // Agent is neither copyable nor movable.
    Agent(const Agent &) = delete;
    Agent &operator=(const Agent &) = delete;

    static Agent *instance() { return gAgent_; }

    void showCollectorControl();

    static QSet<QString> componentSet;
    static QFile recordFile;
    static void initRecordFile();
    static void closeRecordFile();
    static void setAddsValue(const QString &value);
    static bool setContainsValue(const QString &Value);

private slots:
    void onUserEvent(QStringList &);
    void onAppAboutToQuit();

private:
    void onAppStart();
    // 保存数据
    void writeData(QStringList &list);

private:
    static Agent *gAgent_;
    qt_collector::UserEventAnalyzer *eventAnalyzer_ = nullptr;
    QFile dataFile;
    bool openSuccess = false;
    bool isCollecting = true;
    CollectorController controller;

private slots:
    void setCollectionStatus(bool status);
};

} // namespace qt_event_collector


